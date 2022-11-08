CREATE TABLE student
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name  VARCHAR(30) NOT NULL,
    class_id   VARCHAR(10),
    room       VARCHAR(10),
    FOREIGN KEY (class_id) REFERENCES classes (id),
    FOREIGN KEY (room) REFERENCES rooms (room_number)
);

CREATE TABLE rooms
(
    room_number VARCHAR(10) PRIMARY KEY
);

CREATE TABLE classes
(
    id             VARCHAR(10) PRIMARY KEY,
    course         INT         NOT NULL,
    specialization VARCHAR(30) NOT NULL
);

CREATE TABLE record_book
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    id_student INT NOT NULL,
    FOREIGN KEY (id_student) REFERENCES student (id)
        ON DELETE CASCADE
);