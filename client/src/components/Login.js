import React from 'react'
import { useDispatch } from 'react-redux';
import { setUsers } from '../features/users/users';
import { login } from '../api/login';


export default function Login() {
    const dispatch = useDispatch();
    const loginUser = () => {
        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;
        login(username, password).then((res) => {
            const user = res.user;
            const userData = {
                id: user.id,
                username: user.username,
            };
            dispatch(setUsers(userData)); // Dispatch the user data
        }).catch((err) => {
            console.log(err);
        });
    }
  return (
    <div>
        <input id="username"></input>
        <input id="password"></input>
        <button onClick={loginUser}>Login</button>
    </div>
  )
}
