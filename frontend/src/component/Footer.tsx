import { Box, Typography } from "@mui/material";
import React from "react";

import logo from "../asset/logo.png";
import { styles } from "../asset/style";

const customStyle = {
  caption: {
    marginLeft: "4px",
  },
};

export interface IFooterProps {}

const Footer: React.FC<IFooterProps> = () => {
  return (
    <Box sx={styles.footer}>
      <Box sx={styles.footerCaption}>
        <img src={logo} alt="logo" width={16} height={16} />
        <Typography variant="caption" color="initial" sx={customStyle.caption}>
          Copyright ©2022 TutorKit
        </Typography>
      </Box>
    </Box>
  );
};

export default Footer;
