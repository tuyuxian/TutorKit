import { Box } from "@mui/material";
import React from "react";

import CourseCard from "../../component/course/CourseCard";
import CourseCreateCard from "../../component/course/CourseCreateCard";
import { Page } from "../../component/Layout";
export interface ICoursePageProps {}

const CoursePage = () => {
  return (
    <Box sx={{ display: "flex", flexDirection: "row" }}>
      <Page>
        <h3 style={{ display: "flex", marginLeft: "8px" }}>
          Welcome back, Sam 👋
        </h3>
        <Box sx={{ display: "flex", flexWrap: "wrap" }}>
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
