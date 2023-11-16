import React from 'react'
import { useSelector } from 'react-redux';

export default function UserInfo() {
    const user = useSelector(state => state.users.user);
  return (
    <div>
        {
            <div> 
                <h1>Welcome, {user.username}!</h1>
                <p>Courses:</p> 
                </div>
                }
        </div>
  )
}
