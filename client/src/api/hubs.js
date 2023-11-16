import axios from "axios";
const baseUrl = "http://localhost:8080";

///users/:id/hubs
const getHubs = async (userid) => {
    const res = await axios({
        method: "get",
        url: `${baseUrl}/users/${userid}/hubs`,
    });

    const hubs = res.data;
    return hubs
}

export { getHubs };
