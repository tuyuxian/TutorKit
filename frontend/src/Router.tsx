import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import CoursePage from "./page/course";
import ForgotPage from "./page/forgot";
import HomePage from "./page/home";
import LoginPage from "./page/login";
import RegisterPage from "./page/register";

const Router: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/forgot" element={<ForgotPage />} />
        <Route path="/course" element={<CoursePage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
