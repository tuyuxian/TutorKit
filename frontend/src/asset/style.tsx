import { checkboxClasses } from "@mui/material/Checkbox";
import { createTheme } from "@mui/material/styles";

export const theme = createTheme({
  typography: {
    fontFamily: ["Rubik", "sans-serif"].join(","),
    fontSize: 16,
    h1: {
      fontFamily: ["Rubik", "sans-serif"].join(","),
      fontSize: 32,
      color: "#000000",
      fontWeight: "bold",
      verticalAlign: "baseline",
    },
    h2: {
      fontFamily: ["Rubik", "sans-serif"].join(","),
      fontSize: 28,
      letterSpacing: "2px",
      color: "#7B68EE",
      fontWeight: "bold",
      verticalAlign: "baseline",
    },
  },
  components: {
    MuiAppBar: {
      styleOverrides: {
        colorPrimary: {
          backgroundColor: "rgba(240, 240, 252, 0.6)",
        },
      },
    },
    MuiCheckbox: {
      styleOverrides: {
        root: {
          color: "rgb(247,247,247)",
          "&:hover": {
            backgroundColor: "transparent !important",
          },
          [`&.${checkboxClasses.checked}`]: {
            color: "#7B68EE",
          },
          "& .MuiSvgIcon-root": {
            backgroundColor: "rgb(247,247,247)",
            borderRadius: "10px",
            zIndex: 1,
          },
        },
      },
    },
    MuiSpeedDial: {
      styleOverrides: {
        root: {
          "& .MuiFab-primary": { backgroundColor: "#7B68EE", color: "#FFFFFF" },
        },
      },
    },
    MuiListItemButton: {
      styleOverrides: {
        root: {
          transition: "none !important",
          "&.Mui-selected": {
            backgroundColor: "#F0F0FC",
            borderRadius: "0 20px 20px 0",
            "&:hover": {
              backgroundColor: "#F0F0F0",
            },
            color: "#000000",
            "& .MuiListItemIcon-root": {
              color: "#000000",
            },
          },
          ":hover": {
            backgroundColor: "#F0F0F0",
            borderRadius: "0 20px 20px 0",
            color: "#000000",
            "& .MuiListItemIcon-root": {
              color: "#000000",
            },
          },
        },
      },
    },
  },
});

export const styles = {
  topNavHeader: {
    flexGrow: 1,
    marginLeft: "8px",
  },
  topNavButtonBar: {
    width: "-webkit-fill-available",
    zIndex: 1,
    position: "fixed",
    justifyContent: "flex-end",
  },
  sideNavContainer: {
    height: "calc(100% - 64px)",
    width: "inherit",
    position: "absolute",
    display: "flex",
    flexDirection: "column",
    justifyContent: "space-between",
    // backgroundColor: "rgba(240, 240, 252, 0.6)",
  },
  page: {
    display: "flex",
  },
  pageSideBar: {
    width: "220px",
    display: "flex",
    verticalAlign: "top",
  },
  pageContent: {},
  container: {
    display: "flex",
    justifyContent: "center",
  },
  inputArea: {
    flexDirection: "column",
  },
  inputLabel: {
    display: "flex",
    alignItems: "flex-end",
    fontFamily: "Rubik",
    fontSize: 16,
  },
  inputInner: {
    width: "100%",
    padding: "4px 8px 4px 8px",
  },
  inputBar: {
    margin: "8px 0 16px 0",
    display: "flex",
    alignItems: "center",
    width: "100%",
    borderRadius: "10px",
    boxShadow: 0,
    backgroundColor: "rgb(247,247,247)",
  },
  inputBase: {
    width: "100%",
  },
  loginForm: {
    width: "400px",
    height: "560px",
    boxShadow: 0,
    marginTop: "16px",
  },
  registerForm: {
    width: "400px",
    height: "800px",
    boxShadow: 0,
    marginTop: "16px",
  },
  supportText: {
    color: "rgb(118,118,118)",
    fontSize: 16,
  },
  button: {
    boxShadow: 0,
    textTransform: "none",
    fontFamily: "Rubik",
    fontWeight: "bold",
  },
  primaryButton: {
    height: "100%",
    width: "100%",
    marginTop: "16px",
    backgroundColor: "#7B68EE",
    borderRadius: "20px",
    ":hover": { backgroundColor: "#7B68EE", boxShadow: 0 },
  },
  secondaryButton: {
    color: "#7B68EE",
    ":hover": { backgroundColor: "#FFFFFF", boxShadow: 0 },
  },
  footer: {
    bottom: 0,
    width: "100%",
  },
  footerCaption: {
    justifyContent: "center",
    display: "flex",
    alignItems: "center",
    mb: 2,
  },
};
