import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { getCourses, removeCourseFromUser  } from '../api/courses';
import { setCourses, removeCourses } from '../features/courses/courses';
import CourseAdd from './CourseAdd';

export default function Dashboard() {
    const user = useSelector(state => state.users.user);
    const courses = useSelector(state => state.courses.courses);


    //Load courses for user
    const dispatch = useDispatch();
    useEffect(() => {
        getCourses(user.id).then((courses) => {
            dispatch(setCourses(courses));
        });
        console.log("COURSES: " + courses)
    }, []);

    const testRemove = () => {
        const toRemove = {
            courseCode: "CAS CS 210"
        }
        removeCourseFromUser(user.id, toRemove.courseCode).then((res) => {
            dispatch(removeCourses(toRemove));
        }).catch((err) => {
            console.log(err);
        });
        
    }
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
                        <button id="test" onClick={testRemove}>Test course removal</button>
                        <CourseAdd />
                </div>
            ) : "loading"}
        </div>
    );
}
