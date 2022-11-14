CREATE TABLE student
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    first_name    VARCHAR(20) CHECK (first_name != '') NOT NULL,
    last_name     VARCHAR(30) CHECK (last_name != '')  NOT NULL,
    group_id      VARCHAR(10) CHECK (group_id != ''),
    room          VARCHAR(10),
    date_of_birth VARCHAR(10) CHECK (date_of_birth != ''),
    FOREIGN KEY (group_id) REFERENCES `group` (id),
    FOREIGN KEY (room) REFERENCES rooms (room_number)
        ON DELETE SET NULL
);

CREATE TABLE rooms
(
    room_number    VARCHAR(10) CHECK (room_number != '') PRIMARY KEY,
    number_of_beds INT
);

CREATE TABLE `group`
(
    id               VARCHAR(10) CHECK (id != '') PRIMARY KEY,
    course           INT                                      NOT NULL,
    number_of_places INT,
    specialization   VARCHAR(30) CHECK (specialization != '') NOT NULL
);

CREATE TABLE record_book
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    id_student INT NOT NULL,
    FOREIGN KEY (id_student) REFERENCES student (id)
        ON DELETE CASCADE
);



DROP TABLE record_book;
DROP TABLE student;
DROP TABLE rooms;
DROP TABLE `group`