import React from "react";

import { MainButton } from "../../component/HelperButton";
import { Page } from "../../component/Layout";

export interface ICoursePageProps {}

const CoursePage = () => {
  return (
    <div>
      <MainButton currentPage="Course" />
      <Page></Page>
    </div>
  );
};

export default CoursePage;
