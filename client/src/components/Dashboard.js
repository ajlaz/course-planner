import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { getCourses, removeCourseFromUser  } from '../api/courses';
import { setCourses, removeCourses } from '../features/courses/courses';
import CourseAdd from './CourseAdd';
import HubList from './HubList';
import Courses from './Courses';
import UserInfo from './UserInfo';
import SuggestCourses from './SuggestCourses';

export default function Dashboard() {
    const user = useSelector(state => state.users.user);
    const courses = useSelector(state => state.courses.courses);


    //Load courses for user
    const dispatch = useDispatch();
    useEffect(() => {
        getCourses(user.id).then((courses) => {
            dispatch(setCourses(courses));
        });
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
        <div className='dashboard'>
            <div className="courses">
                <UserInfo />
                <Courses />
                <CourseAdd />
                
            </div>
            <HubList />
            <SuggestCourses />
        </div>
    );
}
