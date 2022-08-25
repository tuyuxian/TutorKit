import {
  Box,
  Button,
  Card,
  CardContent,
  FormGroup,
  Typography,
} from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import React from "react";

import { styles, theme } from "../../asset/style";
import Footer from "../../component/Footer";
import { FormCheckbox, FormRow, FormSelect } from "../../component/InputForm";
import TopNavBar from "../../component/TopNavBar";

const customStyle = {
  supportSpan: {
    alignItems: "center",
  },
  formGroup: {
    justifyContent: "space-between",
  },
  phoneBar: {
    display: "flex",
  },
  countryInput: {
    width: "40%",
  },
  phoneNumberInput: {
    flexGrow: 1,
    paddingLeft: "12px",
  },
};

export interface IRegisterPageProps {}

const RegisterPage: React.FC<IRegisterPageProps> = () => {
  return (
    <ThemeProvider theme={theme}>
      <div>
        <TopNavBar isAuth={false} />
        <Box sx={styles.container}>
          <Card sx={styles.registerForm}>
            <CardContent>
              <Typography variant="h1">Sign Up</Typography>
              <span style={{ ...styles.container, ...customStyle.supportSpan }}>
                <Typography sx={styles.supportText}>
                  Already have an account?
                </Typography>
                <Button
                  disableRipple={true}
                  href="/login"
                  sx={{ ...styles.button, ...styles.secondaryButton }}
                >
                  Sign In
                </Button>
              </span>
              <Box sx={styles.inputArea}>
                <FormRow rowHeader="Username" />
                <FormRow rowHeader="Email" />
                <FormRow rowHeader="Password" />
                <Box sx={customStyle.phoneBar}>
                  <Box sx={customStyle.countryInput}>
                    <FormSelect
                      label="Country"
                      selectOptions={
                        <>
                          <option value="US">+1 (US)</option>
                          <option value="TW">+886 (TW)</option>
                        </>
                      }
                    />
                  </Box>
                  <Box sx={customStyle.phoneNumberInput}>
                    <FormRow rowHeader="Phone Number" />
                  </Box>
                </Box>
                <FormRow rowHeader="Date of Birth" />
                <Box>
                  <FormGroup row sx={customStyle.formGroup}>
                    <FormCheckbox label="Tutor" />
                    <FormCheckbox label="Student" />
                    <FormCheckbox label="Parent" />
                  </FormGroup>
                </Box>
              </Box>
              <Box>
                <Button
                  variant="contained"
                  disableRipple={true}
                  sx={{ ...styles.button, ...styles.primaryButton }}
                >
                  Sign Up
                </Button>
              </Box>
            </CardContent>
          </Card>
        </Box>
        <Footer position="fixed" />
      </div>
    </ThemeProvider>
  );
};

export default RegisterPage;
