import { Box } from "@mui/material";
import React from "react";

import { styles } from "../asset/style";
import SideNavBar from "./SideNavBar";

export interface IPageProps {
  children?: React.ReactNode;
}

export const Page: React.FC<IPageProps> = ({ children }) => {
  return (
    <Box sx={styles.page}>
      <Box sx={styles.pageSideBar}>
        <SideNavBar />
      </Box>
      <Box>{children}</Box>
    </Box>
  );
};
