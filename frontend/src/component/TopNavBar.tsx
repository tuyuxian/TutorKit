import AccountCircle from "@mui/icons-material/AccountCircle";
import {
  AppBar,
  Box,
  Button,
  IconButton,
  Link,
  Toolbar,
  Typography,
} from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import React from "react";
import { useMediaQuery } from "react-responsive";
import { useLocation } from "react-router-dom";

import logo from "../asset/logo.png";
import { styles, theme } from "../asset/style";

const customStyle = {
  mobileTopNavButtonBar: {
    position: "relative",
  },
  topNavLogoBar: {
    justifyContent: "center",
  },
  authTopNavLogoBar: {
    justifyContent: "space-between",
  },
  brandLogo: {
    display: "flex",
  },
  link: {
    display: "flex",
    alignItems: "flex-end",
  },
  loginButton: {
    color: "#000000",
    ":hover": { backgroundColor: "inherit", boxShadow: 0 },
    marginRight: "4px",
  },
  registerButton: {
    color: "#FFFFFF",
    backgroundColor: "#7B68EE",
    borderRadius: "20px",
    ":hover": { backgroundColor: "#7B68EE", boxShadow: 0 },
  },
  accountIcon: {
    color: "#7B68EE",
  },
};

export interface ITopNavBarProps {
  isAuth: boolean;
}

const TopNavBar: React.FC<ITopNavBarProps> = ({ isAuth }) => {
  const location = useLocation().pathname;
  const isTabletOrMobile = useMediaQuery({ query: "(max-width: 768px)" });

  return (
    <ThemeProvider theme={theme}>
      <Box>
        <AppBar position="static" elevation={1}>
          {location === "/" && (
            <Toolbar
              sx={
                isTabletOrMobile
                  ? {
                      ...styles.topNavButtonBar,
                      ...customStyle.mobileTopNavButtonBar,
                    }
                  : styles.topNavButtonBar
              }
            >
              <Box>
                <Button
                  disableRipple={true}
                  href="/login"
                  sx={{
                    ...styles.button,
                    ...styles.secondaryButton,
                    ...customStyle.loginButton,
                  }}
                >
                  Sign In
                </Button>
                <Button
                  variant="contained"
                  disableRipple={true}
                  href="/register"
                  sx={{
                    ...styles.button,
                    ...customStyle.registerButton,
                  }}
                >
                  Sign Up
                </Button>
              </Box>
            </Toolbar>
          )}
          <Toolbar
            sx={
              isAuth ? customStyle.authTopNavLogoBar : customStyle.topNavLogoBar
            }
          >
            <Box
              sx={
                isTabletOrMobile && location === "/"
                  ? null
                  : customStyle.brandLogo
              }
            >
              <img src={logo} alt="logo" width={36} height={36} />
              <Link href="/" underline="none" sx={customStyle.link}>
                <Typography
                  variant="h2"
                  component="div"
                  sx={styles.topNavHeader}
                >
                  TutorKit
                </Typography>
              </Link>
            </Box>
            {isAuth && (
              <IconButton size="large" sx={customStyle.accountIcon}>
                <AccountCircle />
              </IconButton>
            )}
          </Toolbar>
        </AppBar>
      </Box>
    </ThemeProvider>
  );
};

export default TopNavBar;
