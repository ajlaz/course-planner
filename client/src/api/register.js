import axios from 'axios';
const baseUrl = 'http://localhost:8080';
const { useSelector } = require("react-redux");

const register = async (username, password) => {
    const res = await axios({
        method: 'post',
        url: `${baseUrl}/register`,
        data: {
            username: username,
            password: password
        }
    })
    return res.data;
}




export {register};
