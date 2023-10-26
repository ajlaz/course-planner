import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { getCourses } from '../api/courses';
import { setCourses } from '../features/courses/courses';
import CourseAdd from './CourseAdd';

export default function Dashboard() {
    const user = useSelector(state => state.users.user);
    const courses = useSelector(state => state.courses.courses.courses);
    const dispatch = useDispatch();
    useEffect(() => {
        getCourses(user.id).then((courses) => {
            dispatch(setCourses(courses));
        });
        console.log(courses)
    }, []);
    return (
        <div>
            {user.id ? (
                <div>
                    <p>User ID: {user.id}</p>
                    <p>Username: {user.username}</p>
                    <p>Courses:</p>
                    <div id="courses"> {
                        courses?.map((course) => (
                            <div key={course.id}>
                                <p>{course.courseCode}</p>
                                <p>{course.courseTitle}</p>
                                <p>{course.hubs}</p>
                            </div>
                            ))
                        }
                        </div>
                        <CourseAdd />
                </div>
            ) : "loading"}
        </div>
    );
}
