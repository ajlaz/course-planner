import React from 'react'
import { getHubs } from '../api/hubs'
import { useSelector } from 'react-redux';
import { useEffect } from 'react';
import { setHubs } from '../features/hubs/hubs';
import { useDispatch } from 'react-redux';


export default function HubList() {
    const hubs = useSelector(state => state.hubs.hubs);
    const user = useSelector(state => state.users.user);
    const courses = useSelector(state => state.courses.courses);
    const dispatch = useDispatch();
    useEffect(() => {
        getHubs(user.id).then((hubs) => {
            dispatch(setHubs(hubs));
        });
    }, [courses]);

    let i = 0;
  return (
    <div>
        HubList
        {
            //iterate through the object hubs and return a p for the key and value
            Object.entries(hubs).map(([key, value]) => {
                return <p key={i++}>{key}: {value}</p>
            })
        }
    </div>
  )
}
