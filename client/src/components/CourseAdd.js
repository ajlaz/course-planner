import React from 'react'
import { useDispatch, useSelector} from 'react-redux';
import { addCourseToUser } from '../api/courses';

export default function CourseAdder() {
    const user = useSelector(state => state.users.user);
    const dispatch = useDispatch();

    const addCourse = () => {
        const courseCode = document.getElementById("courseCode").value;
        addCourseToUser(user.id, courseCode).then((res) => {
            console.log(res);
        }).catch((err) => {
            console.log(err);
        });
    }
    return (
        <div>
            <input id="courseCode"></input>
            <button onClick={addCourse}>Add Course</button>
        </div>
    )
}

