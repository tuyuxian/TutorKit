import {
  AssignmentTurnedIn,
  Forum,
  Home,
  LogoutOutlined,
  SettingsOutlined,
  ViewTimeline,
} from "@mui/icons-material";
import {
  Avatar,
  Box,
  Button,
  Divider,
  Link,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
} from "@mui/material/";
import { ThemeProvider } from "@mui/material/styles";
import React, { useState } from "react";

import logo from "../asset/logo.png";
import { styles, theme } from "../asset/style";

const customStyle = {
  flex: {
    display: "flex",
    justifyContent: "center",
  },
  brandLogo: {
    margin: "16px",
  },
  link: {
    display: "flex",
    alignItems: "center",
  },
  listIcon: {
    marginLeft: "20px",
  },
  listItem: {
    margin: "30px 0 30px 60px",
    height: "48px",
  },
  backgroundDiv: {
    height: "calc(100%)",
    width: "285px",
    margin: "0 0 0 0",
    borderRadius: "0 32.8px 32.8px 0",
    backgroundColor: "#E1E3E4",
    boxShadow: "5px 3px 10px rgba(0,0,0,0.2)",
  },
  sideBarButton: {
    height: "40px",
    width: "220px",
    border: "2.5px solid #E1E3E4",
    borderRadius: "11px",
    margin: "4px 0px 4px 0px",
  },
  buttonGroup: {
    display: "flex",
    flexDirection: "column",
    margin: "16px 30px 16px 30px",
  },
  divider: {
    width: "220px",
    borderBottomWidth: "2.5px",
    borderRadius: "11px",
    color: "#E1E3E4",
  },
  accountDiv: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
  },
  accountInfo: {
    display: "flex",
    marginLeft: "8px",
  },
};

const sideMenuList = [
  {
    listIcon: <Home />,
    listText: "Course",
  },
  {
    listIcon: <ViewTimeline />,
    listText: "Todo",
  },
  {
    listIcon: <AssignmentTurnedIn />,
    listText: "Attendence",
  },
  {
    listIcon: <Forum />,
    listText: "Community",
  },
];

export interface ISideNavBarProps {}

const SideNavBar = () => {
  const [selectedIndex, setSelectedIndex] = useState<number>(0);
  const handleListItemClick = (index: number) => {
    setSelectedIndex(index);
  };
  return (
    <ThemeProvider theme={theme}>
      <Box
        component="div"
        sx={{ ...styles.sideNavContainer, ...customStyle.backgroundDiv }}
      ></Box>
      <Box component="div" sx={styles.sideNavContainer}>
        <Box>
          <Box sx={{ ...customStyle.flex, ...customStyle.brandLogo }}>
            <img src={logo} alt="logo" width={48} height={48} />
            <Link href="/" underline="none" sx={customStyle.link}>
              <Typography variant="h2" component="div" sx={styles.topNavHeader}>
                TutorKit
              </Typography>
            </Link>
          </Box>
          <List>
            {sideMenuList.map((listItem, index) => (
              <ListItemButton
                disableRipple
                key={index}
                sx={customStyle.listItem}
                selected={selectedIndex === index}
                onClick={() => handleListItemClick(index)}
              >
                <ListItemIcon sx={customStyle.listIcon}>
                  {listItem.listIcon}
                </ListItemIcon>
                <ListItemText>
                  <Typography variant="h1" sx={styles.inputLabel}>
                    {listItem.listText}
                  </Typography>
                </ListItemText>
              </ListItemButton>
            ))}
          </List>
        </Box>
        <Box>
          <Box sx={customStyle.flex}>
            <Divider variant="fullWidth" sx={customStyle.divider} />
          </Box>
          <Box sx={customStyle.buttonGroup}>
            <Box sx={customStyle.accountDiv}>
              <Avatar>T</Avatar>
              <Box>
                <Typography
                  variant="subtitle1"
                  component="div"
                  color="secondary"
                  sx={customStyle.accountInfo}
                >
                  Test User
                </Typography>
                <Typography
                  variant="body2"
                  component="div"
                  sx={customStyle.accountInfo}
                >
                  testuser@gmail.com
                </Typography>
              </Box>
            </Box>
          </Box>
          <Box sx={customStyle.buttonGroup}>
            <Button
              variant="outlined"
              disableRipple={true}
              color="secondary"
              startIcon={<SettingsOutlined />}
              sx={customStyle.sideBarButton}
            >
              Settings
            </Button>
            <Button
              variant="outlined"
              disableRipple={true}
              color="secondary"
              startIcon={<LogoutOutlined />}
              sx={customStyle.sideBarButton}
            >
              Logout
            </Button>
          </Box>
        </Box>
      </Box>
    </ThemeProvider>
  );
};

export default SideNavBar;
