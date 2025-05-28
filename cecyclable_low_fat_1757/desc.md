+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| product_id  | int     |
| low_fats    | enum    |
| recyclable  | enum    |
+-------------+---------+
product_id is the primary key (column with unique values) for this table.
low_fats is an ENUM (category) of type ('Y', 'N') where 'Y' means this product is low fat and 'N' means it is not.
recyclable is an ENUM (category) of types ('Y', 'N') where 'Y' means this product is recyclable and 'N' means it is not.

| student_id | student_name | subject_name | student_id | subject_name | count(e.student_id) | e.student_id is not null |
| ---------- | ------------ | ------------ | ---------- | ------------ | ------------------- | ------------------------ |
| 1          | Alice        | Programming  | 1          | Programming  | 1                   | 1                        |
| 1          | Alice        | Physics      | 1          | Physics      | 2                   | 1                        |
| 1          | Alice        | Math         | 1          | Math         | 3                   | 1                        |
| 2          | Bob          | Programming  | 2          | Programming  | 1                   | 1                        |
| 2          | Bob          | Physics      | null       | null         | 0                   | 0                        |
| 2          | Bob          | Math         | 2          | Math         | 1                   | 1                        |
| 13         | John         | Programming  | 13         | Programming  | 1                   | 1                        |
| 13         | John         | Physics      | 13         | Physics      | 1                   | 1                        |
| 13         | John         | Math         | 13         | Math         | 1                   | 1                        |
| 6          | Alex         | Programming  | null       | null         | 0                   | 0                        |

| student_id | student_name | subject_name | student_id | subject_name | e.student_id is not null |
| ---------- | ------------ | ------------ | ---------- | ------------ | ------------------------ |
| 1          | Alice        | Programming  | 1          | Programming  | 1                        |
| 1          | Alice        | Physics      | 1          | Physics      | 1                        |
| 1          | Alice        | Physics      | 1          | Physics      | 1                        |
| 1          | Alice        | Math         | 1          | Math         | 1                        |
| 1          | Alice        | Math         | 1          | Math         | 1                        |
| 1          | Alice        | Math         | 1          | Math         | 1                        |
| 2          | Bob          | Programming  | 2          | Programming  | 1                        |
| 2          | Bob          | Physics      | null       | null         | 0                        |
| 2          | Bob          | Math         | 2          | Math         | 1                        |
| 13         | John         | Programming  | 13         | Programming  | 1                        |
| 13         | John         | Physics      | 13         | Physics      | 1                        |
| 13         | John         | Math         | 13         | Math         | 1                        |
| 6          | Alex         | Programming  | null       | null         | 0                        |
| 6          | Alex         | Physics      | null       | null         | 0                        |
| 6          | Alex         | Math         | null       | null         | 0                        |

| student_id | student_name | subject_name | student_id | subject_name | count(e.student_id) | e.student_id is not null |
| ---------- | ------------ | ------------ | ---------- | ------------ | ------------------- | ------------------------ |
| 1          | Alice        | Programming  | 1          | Programming  | 1                   | 1                        |
| 1          | Alice        | Physics      | 1          | Physics      | 2                   | 1                        |
| 1          | Alice        | Math         | 1          | Math         | 3                   | 1                        |
| 2          | Bob          | Programming  | 2          | Programming  | 1                   | 1                        |
| 2          | Bob          | Physics      | null       | null         | 0                   | 0                        |
| 2          | Bob          | Math         | 2          | Math         | 1                   | 1                        |
| 13         | John         | Programming  | 13         | Programming  | 1                   | 1                        |
| 13         | John         | Physics      | 13         | Physics      | 1                   | 1                        |
| 13         | John         | Math         | 13         | Math         | 1                   | 1                        |
| 6          | Alex         | Programming  | null       | null         | 0                   | 0                        |
| 6          | Alex         | Physics      | null       | null         | 0                   | 0                        |
| 6          | Alex         | Math         | null       | null         | 0                   | 0                        |