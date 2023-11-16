import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { getCourses, removeCourseFromUser  } from '../api/courses';
import { setCourses, removeCourses } from '../features/courses/courses';

export default function Courses() {
    const user = useSelector(state => state.users.user);
    const courses = useSelector(state => state.courses.courses);


    //Load courses for user
    const dispatch = useDispatch();
    useEffect(() => {
        getCourses(user.id).then((courses) => {
            dispatch(setCourses(courses));
        });
    }, []);

    const remove = (toRemove) => {
        removeCourseFromUser(user.id, toRemove.courseCode).then((res) => {
            dispatch(removeCourses(toRemove));
        }).catch((err) => {
        });
        
    }
    return (
        <div>
            {user.id ? (
                <div>
                    <div id="courses"> {
                        courses?.map((course) => (
                            <div key={course.id} className='course'>
                                <p>{course.courseCode}</p>
                                <button onClick={() => remove(course)} >X</button>
                            </div>
                            
                            ))
                        }
                        </div>
                </div>
            ) : "loading"}
        </div>
    );
}
