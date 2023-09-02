const fs = require("fs");
const path = require("path");
const mysql = require("mysql2/promise");
const dotenv = require("dotenv");
dotenv.config();

const dataFolderPath = process.env.DATA_FOLDER_PATH || "./data/";

const githubUserName = process.env.GITHUB_USERNAME || "debasishbsws";
const githubRepoName = process.env.GITHUB_REPO || "question-book-GO";


const InstituteTableDATA = [];
const SubjectTableDATA = [];
const InstituteSubjectTableDATA = [];
const QuestionPapersTableDATA = [];


// Creating Table details from the /data folder

class TreeNode {
    constructor(name, fullPath, level, parent) {
        this.name = name;
        this.fullPath = fullPath;
        this.children = [];
        this.level = level;
        this.parent = parent;
    }
}

function buildTreeStructure(folderPath, level) {
    const rootNode = new TreeNode("data", folderPath, level, null);
    traverseFolders(folderPath, rootNode);
    return rootNode;
}

function traverseFolders(folderPath, parentNode) {
    const files = fs.readdirSync(folderPath);
    for (const file of files) {
        const fullPath = path.join(folderPath, file);
        const stats = fs.statSync(fullPath);

        const newNode = new TreeNode(file, fullPath, parentNode.level + 1, parentNode);
        parentNode.children.push(newNode);
        if (stats.isDirectory()) {
            traverseFolders(fullPath, newNode);
        } 
    }
}

function BFSTraversAndFillDetails(rootNode) {
    const queue = [rootNode];
    while (queue.length > 0) {
        const node = queue.shift();
        for (const child of node.children) {
            queue.push(child);
        }
        const stats = fs.statSync(node.fullPath);
        if(node.level === 2){
            if(stats.isDirectory() && !node.name.endsWith(".json")) {
                // TODO: check if subject exists in subjectDetails
                const instituteSubject = {
                    subject_id: node.name,
                    institute_id: node.parent.name
                };
                InstituteSubjectTableDATA.push(instituteSubject);
            }
            else if (stats.isFile() && node.name ==="institute_details.json") {
                addInstituteDetails(node.fullPath, node.parent);
            }
        }
        if(stats.isFile() && node.name.endsWith(".pdf")) {
            addQuestionPaperDetailsFromPath(node.fullPath);
        }
    }
}

function addInstituteDetails(instituteJsonFullPath, parentNode) {
    const instituteDetailsFromJson = JSON.parse(fs.readFileSync(instituteJsonFullPath));
    instituteDetailsFromJson["id"] = parentNode.name;
    InstituteTableDATA.push(instituteDetailsFromJson);
}

function getSubjectDetailsFromFile() {
    const subjectDetailsJson = JSON.parse(fs.readFileSync(path.join(dataFolderPath,"subjects.json")));
    for (const subject of subjectDetailsJson) {
        SubjectTableDATA.push(subject);
    }
}

function addQuestionPaperDetailsFromPath(paperPath) {
    const parsedPath = path.parse(paperPath);
    const dirList = parsedPath.dir.split(path.sep);
    // remove the first data folder
    dirList.shift();
    let [institute, subject, year, semester] = dirList;
    let [exam_type, title] = parsedPath.name.split("-");
    title = capitalizeWords(title);
    semester = capitalizeWords(semester);
    exam_type = capitalizeWords(exam_type);


    const url = createUrlFromPath(paperPath);

    QuestionPapersTableDATA.push({
        url,
        year,
        semester,
        exam_type,
        title,
        institute,
        subject
    });
}
// util functions
function createUrlFromPath(paperPath) {
    const url = paperPath.replace(/\\/g, "/");
    return githubUserName+"/"+ githubRepoName +"/main/" + url;
}

function capitalizeWords(str) {
    return str
        .split("_")
        .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
        .join(" ");
}

const root = buildTreeStructure(dataFolderPath, 0);
BFSTraversAndFillDetails(root);
getSubjectDetailsFromFile();

//-------------------------------------------------
//  DB functions and update the database according to the data
//-------------------------------------------------

// const UserName = process.env.DB_USER_NAME || "dev";
// const Password = process.env.DB_PASSWORD || "password" ;
// const Database = process.env.DATABASE_NAME || "devDB";
// const Host = process.env.DB_HOST || "localhost";

// const pool = mysql.createPool({
//     host: Host,
//     user: UserName,
//     password: Password,
//     database: Database,
//     waitForConnections: true,
//     connectionLimit: 10,
//     queueLimit: 0
// });

const DATABASE_URI = process.env.DATABASE_URI;
const pool = mysql.createPool(DATABASE_URI);

async function executeQuery (sql, values) {
    let connection;
    try {
        let connection = await pool.getConnection();
        const [rows] = await connection.query(sql, values);
        connection.release();
        return rows;
    } finally {
        if (connection != null) {
            console.log("Connection released finally");
            connection.release();
        }
    }
}

async function UpdateInstituteTable(){
    console.log("Updating Institute Table Started");
    await executeQuery("SELECT * FROM institute").then(async (institutes) => {
        const InstituteTableDB = institutes;

        for( const institute of InstituteTableDATA){
        // check if it is already present in InstituteFromDB
            const instituteDB =InstituteTableDB.find((ins) => ins.id === institute.id);
            if(instituteDB === undefined){
                const query = "INSERT INTO institute (id, institute_name, alt_name, website, country, state) VALUES (?, ?, ?, ?, ?, ?)";
                const altNameJSON = JSON.stringify(institute.altName);
                const values = [institute.id, institute.name, altNameJSON, institute.website, institute.country, institute.state];
                await executeQuery(query, values);
            }

            // check if any institutes details are changed
            else if(instituteDB.institute_name !== institute.name || instituteDB.website !== institute.website || instituteDB.country !== institute.country || instituteDB.state !== institute.state){
                const query = "UPDATE institute SET institute_name = ?, website = ?, country = ?, state = ? WHERE id = ?";
                const values = [institute.name, institute.website, institute.country, institute.state, institute.id];
                await executeQuery(query, values);
            }
        }
        // check if any institute is deleted
        // TODO: Delete might not work as it is referenced in institute_subject table fk constraint
        for(const instituteDB of InstituteTableDB){
            const institute = InstituteTableDATA.find((institute) => institute.id === instituteDB.id);
            if(institute === undefined){
                const query = "DELETE FROM institute WHERE id = ?";
                const values = [instituteDB.id];
                await executeQuery(query, values);
            }
        }
    });
}



async function UpdateSubjectTable(){
    console.log("Updating Subject Table Started");
    await executeQuery("SELECT * FROM subject").then(async (subjects) => {
        const SubjectTableDB = subjects;

        for( const subject of SubjectTableDATA){
        // check if it is already present in subjectFromDB
            const subjectDB = SubjectTableDB.find((sub) => sub.id === subject.id);
            if(subjectDB === undefined){
                const query = "INSERT INTO subject (id, subject_name, synonyms) VALUES (?, ?, ?)";
                const synonymsJSON = JSON.stringify(subject.synonyms);
                const values = [subject.id, subject.subject_name, synonymsJSON];
                await executeQuery(query, values);
            }

            // check if any subjects details are changed
            else if(subjectDB.subject_name !== subject.subject_name){
                const query = "UPDATE subject SET subject_name = ? WHERE id = ?";
                const values = [subject.subject_name, subject.id];
                await executeQuery(query, values);
            }
        }
        // check if any subject is deleted
        // TODO: Delete might not work as it is referenced in institute_subject table fk constraint
        for(const subjectDB of SubjectTableDB){
            const subject = SubjectTableDATA.find((subject) => subject.id === subjectDB.id);
            if(subject === undefined){
                const query = "DELETE FROM subject WHERE id = ?";
                const values = [subjectDB.id];
                await executeQuery(query, values);
            }
        }
    });
}


async function UpdateInstituteSubjectTable(){
    console.log("Updating Institute Subject Table Started");
    executeQuery("SELECT * FROM institute_subject").then((instituteSubjects) => {
        const InstituteSubjectTableDB = instituteSubjects;

        for( const instituteSubject of InstituteSubjectTableDATA){
        // check if it is already present in instituteSubjectFromDB
            const instituteSubjectDB = InstituteSubjectTableDB.find((insSub) => insSub.institute_id === instituteSubject.institute_id && insSub.subject_id === instituteSubject.subject_id);
            if(instituteSubjectDB === undefined){
                const query = "INSERT INTO institute_subject (institute_id, subject_id) VALUES (?, ?)";
                const values = [instituteSubject.institute_id, instituteSubject.subject_id];
                executeQuery(query, values);
            }
        }
        // check if any institute_subject is deleted
        for(const instituteSubjectDB of InstituteSubjectTableDB){
            const instituteSubject = InstituteSubjectTableDATA.find((instituteSubject) => instituteSubject.institute_id === instituteSubjectDB.institute_id && instituteSubject.subject_id === instituteSubjectDB.subject_id);
            if(instituteSubject === undefined){
                const query = "DELETE FROM institute_subject WHERE institute_id = ? AND subject_id = ?";
                const values = [instituteSubjectDB.institute_id, instituteSubjectDB.subject_id];
                executeQuery(query, values);
            }
        }
    });
}


async function UpdateQuestionPaperTable(){
    console.log("Updating Question Paper Table Started");
    executeQuery("SELECT * FROM question_paper").then((questionPapers) => {
        const QuestionPaperTableDB = questionPapers;

        for( const questionPaper of QuestionPapersTableDATA){
        // check if it is already present in questionPaperFromDB
            const questionPaperDB = QuestionPaperTableDB.find((qp) => qp.url === questionPaper.url);
            if(questionPaperDB === undefined){
                const query = "INSERT INTO question_paper (url, year, semester, exam_type, title, institute_id, subject_id) VALUES (?, ?, ?, ?, ?, ?, ?)";
                const values = [questionPaper.url, questionPaper.year, questionPaper.semester, questionPaper.exam_type, questionPaper.title, questionPaper.institute, questionPaper.subject];
                executeQuery(query, values);
            }
        }
        // check if any question paper is deleted
        for(const questionPaperDB of QuestionPaperTableDB){
            const questionPaper = QuestionPapersTableDATA.find((questionPaper) => questionPaper.url === questionPaperDB.url);
            if(questionPaper === undefined){
                const query = "DELETE FROM question_paper WHERE url = ?";
                const values = [questionPaperDB.url];
                executeQuery(query, values);
            }
        }
    });
}

async function UpdateDatabase(){
    await UpdateInstituteTable();
    await UpdateSubjectTable();
    await UpdateInstituteSubjectTable();
    await UpdateQuestionPaperTable();
}

UpdateDatabase().then(() => {
    console.log("Done");
});

setTimeout(() => {
    process.exit(0);
}, 10000);