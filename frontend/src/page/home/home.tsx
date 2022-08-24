import { Box } from "@mui/material";
import React from "react";

import Footer from "../../component/Footer";
import TopNavBar from "../../component/TopNavBar";

export interface ILoginPageProps {}

const HomePage: React.FC<ILoginPageProps> = () => {
  return (
    <div>
      <TopNavBar />
      <Box>This is home page</Box>
      <Footer />
    </div>
  );
};

export default HomePage;
