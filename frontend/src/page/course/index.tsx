import { Box } from "@mui/material";
import React from "react";

import CourseCard from "../../component/course/CourseCard";
import CourseCreateCard from "../../component/course/CourseCreateCard";
import { Page } from "../../component/Layout";

const customStyle = {
  header: {
    display: "flex",
    marginLeft: "8px",
    fontFamily: "Poppins",
  },
  main: {
    display: "flex",
    flexDirection: "row",
  },
  content: {
    display: "flex",
    flexWrap: "wrap",
  },
};
export interface ICoursePageProps {}

const CoursePage = () => {
  return (
    <Box sx={customStyle.main}>
      <Page>
        <h3 style={customStyle.header}>Welcome back, Sam 👋</h3>
        <Box sx={customStyle.content}>
          <CourseCard usage="tutor" />
          <CourseCard usage="tutor" />
          <CourseCard usage="parents" />
          <CourseCard usage="student" />
          <CourseCard usage="tutor" />
          <CourseCard usage="student" />
          <CourseCreateCard />
        </Box>
      </Page>
    </Box>
  );
};

export default CoursePage;
