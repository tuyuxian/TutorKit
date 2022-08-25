import React from "react";

import { MainButton } from "../../component/HelperButton";
import { Page } from "../../component/Layout";
import TopNavBar from "../../component/TopNavBar";

export interface ICoursePageProps {}

const CoursePage = () => {
  return (
    <div>
      <TopNavBar isAuth={true} />
      <Page></Page>
      <MainButton currentPage="Course" />
    </div>
  );
};

export default CoursePage;
