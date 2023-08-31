USE devDB;

-- Uncomment the following lines before docker compose to resetup the database

-- DROP TABLE IF EXISTS `institute`;
-- DROP TABLE IF EXISTS `subject`;
-- DROP TABLE IF EXISTS `institute_subject`;
-- DROP TABLE IF EXISTS `question_paper`;

-- Create institute table
CREATE TABLE `institute` (
  `id` varchar(255) PRIMARY KEY,
  `institute_name` varchar(255) NOT NULL,
  `alt_name` JSON,
  `website` text,
  `country` varchar(255),
  `state` varchar(255)
);

-- Create subject table
CREATE TABLE `subject` (
  `id` varchar(255) PRIMARY KEY,
  `subject_name` varchar(255),
  `synonyms` JSON
);


-- Create institute_subject table with DELETE on CASCADE
CREATE TABLE `institute_subject` (
  `institute_id` varchar(255),
  `subject_id` varchar(255),
  PRIMARY KEY (`institute_id`, `subject_id`),
  FOREIGN KEY (`institute_id`) REFERENCES `institute` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`subject_id`) REFERENCES `subject` (`id`) ON DELETE CASCADE
);


-- Create question_paper table with DELETE on CASCADE
CREATE TABLE `question_paper` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `url` text,
  `title` varchar(255),
  `year` varchar(4),
  `semester` varchar(10),
  `exam_type` varchar(10),
  `institute_id` varchar(255),
  `subject_id` varchar(255),
  FOREIGN KEY (`institute_id`) REFERENCES `institute` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`subject_id`) REFERENCES `subject` (`id`) ON DELETE CASCADE
);

-- -- Populate `institute` table
-- INSERT INTO `institute` (`id`, `institute_name`, `alt_name`, `website`, `country`, `state`)
-- VALUES ('demo_institute','Demo Institute', '["demo-alt_name1", "demo-alt_name2", "demo-alt_name3"]', 'https://www.demo.website.com/', 'demo-country', 'demo-state'),
--        ('national_institute_of_technology_jamshedpur','National Institute of Technology Jamshedpur', '["NIT Jamshedpur", "NITJSR", "NIT Jam"]', 'https://www.nitjsr.ac.in/', 'India', 'Jharkhand');

-- -- Populate `subject` table
-- INSERT INTO `subject` (`id`, `subject_name`, `synonyms`)
-- VALUES ('demo_subject_1', 'Demo Subject 1', '["demosubject1 Alt 1", "demosubject1 Alt 2"]'),
--        ('demo_subject_2', 'Demo Subject 2', '["demosubject2 Alt 1", "demosubject2 Alt 2"]'),
--        ('demo_subject_3', 'Demo Subject 3', '["demosubject3 Alt 1", "demosubject3 Alt 2"]');

-- -- Populate `institute_subject` table
-- INSERT INTO `institute_subject` (`institute_id`, `subject_id`)
-- VALUES ('demo_institute', 'demo_subject_1'),
--        ('demo_institute', 'demo_subject_2');

-- -- Populate `question_paper` table
-- INSERT INTO `question_paper` (`url`, `year`, `semester`, `exam_type`, `title`, `institute_id`, `subject_id`)
-- VALUES ('https://github.com/debasishbsws/question-book/blob/main/data/demo_institute/demo_subject_1/2022/spring/ENDSEM-Demo_Question_Paper_Title_1.pdf', '2022', 'spring', 'ENDSEM', 'DEMO Question Paper Title 1', 'demo_institute', 'demo_subject_1'),
--        ('https://github.com/debasishbsws/question-book/blob/main/data/demo_institute/demo_subject_1/2022/spring/MIDSEM-Demo_Question_Paper_Title_1.pdf', '2022', 'spring', 'MIDSEM', 'DEMO Question Paper Title 1', 'demo_institute', 'demo_subject_1'),
--        ('https://github.com/debasishbsws/question-book/blob/main/data/demo_institute/demo_subject_1/2022/spring/MIDSEM-Demo_Question_Paper_Title_2.pdf', '2022', 'spring', 'MIDSEM', 'DEMO Question Paper Title 2', 'demo_institute', 'demo_subject_1'),
--        ('https://github.com/debasishbsws/question-book/blob/main/data/demo_institute/demo_subject_1/2022/autumn/ENDSEM-Demo_Question_Paper_Title_3.pdf', '2022', 'autumn', 'ENDSEM', 'DEMO Question Paper Title 3', 'demo_institute', 'demo_subject_1'),
--        ('https://github.com/debasishbsws/question-book/blob/main/data/demo_institute/demo_subject_1/2021/autumn/ENDSEM-Demo_Question_Paper_Title_1.pdf', '2021', 'autumn', 'ENDSEM', 'DEMO Question Paper Title 1', 'demo_institute', 'demo_subject_1'),
--        ('https://github.com/debasishbsws/question-book/blob/main/data/demo_institute/demo_subject_2/2021/spring/ENDSEM-Demo_Question_Paper_Title_4.pdf', '2021', 'spring', 'ENDSEM', 'DEMO Question Paper Title 4', 'demo_institute', 'demo_subject_2');
