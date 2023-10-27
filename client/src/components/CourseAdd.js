import React, { useEffect } from 'react'
import { useDispatch, useSelector} from 'react-redux';
import { addCourseToUser, getCourseById } from '../api/courses';
import {addCourses } from '../features/courses/courses';

export default function CourseAdder() {
    const user = useSelector(state => state.users.user);
    const courses = useSelector(state => state.courses.courses);
    const dispatch = useDispatch();

    const addCourse = () => {
        const courseCode = document.getElementById("courseCode").value;
        addCourseToUser(user.id, courseCode).then((res) => {
            console.log(res);
        }).catch((err) => {
            console.log(err);
        });
        getCourseById(courseCode).then((course) => {
            const inCourses = courses.some((temp) => temp.courseCode === courseCode);
            if(!inCourses){
                dispatch(addCourses(course));
            }
        });

    }

    useEffect(() => {
        
    }, [])
    return (
        <div>
            <input id="courseCode"></input>
            <button onClick={addCourse}>Add Course</button>
        </div>
    )
}