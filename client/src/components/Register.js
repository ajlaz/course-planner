import React from 'react'
import { register } from '../api/register';

export default function Register() {
    const registerUser = () => {
        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;
        register(username, password).then((res) => {
            document.getElementById("reg-status").innerHTML = "registration successful. return to login";
        }).catch((err) => {
            document.getElementById("reg-status").innerHTML = "registration failed, user already exists";
        });
    }
  return (
    <div className="register">
        <input id="username"></input>
        <input id="password"></input>
        <button onClick={registerUser}>Register</button>
        <p id="reg-status"></p>
    </div>
  )
}
