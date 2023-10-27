import { useSelector } from "react-redux";
import axios from "axios";
import { useUserId } from "./login";
const baseUrl = "http://localhost:8080";

const getCourses = async (userid) => {
    const res = await axios({
        method: "get",
        url: `${baseUrl}/users/${userid}/courses`,
    });

    const courses = res.data;
    console.log(courses)
    return courses;
}

const getCourseById = async (courseid) => {
    const res = await axios({
        method: "get",
        url: `${baseUrl}/courses/${courseid}`,
    })
    return res.data;
}

const removeCourseFromUser = async (userid, courseid) => {
    const res = await axios({
        method: "delete",
        url: `${baseUrl}/users/courses`,
        data: {
            userID: userid,
            courseCodes: courseid
        }
    })
    return res
}
/*

func (a *API) GetCourse(c *gin.Context) {
	courseId := c.Param("code")
	course := postgres.Select(courseId, a.db)
	c.JSON(200, gin.H{"course": course})
}
*/

const addCourseToUser = async (userid, courseid) => {
    const res = await axios({
        method: "post",
        url: `${baseUrl}/users/courses`,
        data: {
            userID: userid,
            courseCodes: courseid
        }
    })
    return res
}

export { addCourseToUser, getCourses, getCourseById, removeCourseFromUser};