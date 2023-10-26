import axios from 'axios';
const baseUrl = 'http://localhost:8080';
const { useSelector } = require("react-redux");

const login = async (username, password) => {
    const res = await axios({
        method: 'post',
        url: `${baseUrl}/login`,
        data: {
            username: username,
            password: password
        }
    })
    return res.data;
}





export {login};
