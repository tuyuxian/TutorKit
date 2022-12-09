import { checkboxClasses } from "@mui/material/Checkbox";
import { createTheme } from "@mui/material/styles";

export const theme = createTheme({
  palette: {
    secondary: {
      main: "#7B68EE",
    },
  },
  typography: {
    fontFamily: ["Poppins", "sans-serif"].join(","),
    fontSize: 16,
    h1: {
      fontFamily: ["Poppins", "sans-serif"].join(","),
      fontSize: 32,
      color: "#000000",
      fontWeight: "bold",
      verticalAlign: "baseline",
    },
    h2: {
      fontFamily: ["Poppins", "sans-serif"].join(","),
      fontSize: 28,
      letterSpacing: "1px",
      color: "#7B68EE",
      fontWeight: "bold",
      verticalAlign: "baseline",
    },
    button: {
      textTransform: "none",
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
    MuiCardHeader: {
      styleOverrides: {
        root: {
          "& .MuiCardHeader-content": {
            display: "flex",
            flexDirection: "column",
            alignItems: "flex-start",
            width: "100%",
          },
        },
      },
    },
    MuiListItemButton: {
      styleOverrides: {
        root: {
          transition: "none !important",
          "&.Mui-selected": {
            backgroundColor: "#E1E3E4",
            borderRadius: "48px 0 0 48px",
            height: "48px",
            "&::before": {
              content: '""',
              position: "absolute",
              backgroundColor: "none",
              bottom: "48px",
              left: "195px",
              height: "50px",
              width: "25px",
              borderRadius: "0 0 25px 0",
              boxShadow: "0 20px 0 0 #E1E3E4",
              pointerEvents: "none",
            },
            "&::after": {
              content: '""',
              position: "absolute",
              backgroundColor: "none",
              bottom: "-50px",
              left: "195px",
              height: "50px",
              width: "25px",
              borderRadius: "0 25px 0 0",
              boxShadow: "0 -20px 0 0 #E1E3E4",
              pointerEvents: "none",
            },
            "&:hover": {
              backgroundColor: "#E1E3E4",
              borderRadius: "48px 0 0 48px",
              height: "48px",
            },
            color: "#000000",
            "& .MuiListItemIcon-root": {
              color: "#000000",
            },
          },
          ":hover": {
            backgroundColor: "#E1E3E4",
            borderRadius: "48px 0 0 48px",
            height: "48px",
            color: "#000000",
            "& .MuiListItemIcon-root": {
              color: "#000000",
            },
            "&::before": {
              content: '""',
              position: "absolute",
              backgroundColor: "none",
              bottom: "48px",
              left: "195px",
              height: "50px",
              width: "25px",
              borderRadius: "0 0 25px 0",
              boxShadow: "0 20px 0 0 #E1E3E4",
              pointerEvents: "none",
            },
            "&::after": {
              content: '""',
              position: "absolute",
              backgroundColor: "none",
              bottom: "-50px",
              left: "195px",
              height: "50px",
              width: "25px",
              borderRadius: "0 25px 0 0",
              boxShadow: "0 -20px 0 0 #E1E3E4",
              pointerEvents: "none",
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
    zIndex: 100,
    height: "calc(100% - 10px)",
    width: "280px",
    position: "fixed",
    display: "flex",
    flexDirection: "column",
    justifyContent: "space-between",
    backgroundColor: "#F2F2F4",
    borderRadius: "0 27.8px 27.8px 0",
    margin: "5px 0 0 0",
  },
  page: {
    display: "flex",
  },
  pageSideBar: {
    display: "flex",
    flex: "0 0 300px",
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
    fontFamily: "Poppins",
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
    fontFamily: "Poppins",
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
