CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    password_hash varchar(255) not null,
    email         varchar(255) not null unique,
    role          varchar(255) not null
);

CREATE TABLE surveys
(
    id        serial       not null unique,
    types      varchar(255) not null
);

CREATE TABLE users_surveys
(
    id        serial                                        not null unique,
    user_id   int references users (id) on delete cascade   not null,
    survey_id int references surveys (id) on delete cascade not null
);

CREATE TABLE questions
(
    id          serial       not null unique,
    description varchar(255) not null
);


CREATE TABLE surveys_questions
(
    id          serial                                          not null unique,
    question_id int references questions (id) on delete cascade not null,
    survey_id   int references surveys (id) on delete cascade   not null
);

CREATE TABLE answers
(
    id          serial       not null unique,
    description varchar(255) not null,
    amount      int          not null default 0
);

CREATE TABLE questions_answers
(
    id          serial                                          not null unique,
    question_id int references questions (id) on delete cascade not null,
    answer_id   int references answers (id) on delete cascade   not null
);