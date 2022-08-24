import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import ForgotPage from "./page/forgot/forgot";
import HomePage from "./page/home/home";
import LoginPage from "./page/login/login";
import RegisterPage from "./page/register/register";

const Router: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/forgot" element={<ForgotPage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
