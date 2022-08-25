import { Box, Button, Card, CardContent, Typography } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import React from "react";

import { styles, theme } from "../../asset/style";
import Footer from "../../component/Footer";
import { FormRow } from "../../component/InputForm";
import TopNavBar from "../../component/TopNavBar";

const customStyle = {
  email: {
    marginTop: "32px",
  },
  backToLoginButton: {
    color: "#000000",
    backgroundColor: "#F0F0FC",
    ":hover": { backgroundColor: "#F0F0FC", boxShadow: 0 },
  },
};

export interface IForgotPageProps {}

const ForgotPage: React.FC<IForgotPageProps> = () => {
  return (
    <ThemeProvider theme={theme}>
      <div>
        <TopNavBar isAuth={false} />
        <Box sx={styles.container}>
          <Card sx={styles.loginForm}>
            <CardContent>
              <Typography variant="h1">Forgot Password?</Typography>
              <Box sx={{ ...styles.inputArea, ...customStyle.email }}>
                <FormRow rowHeader="Email" />
              </Box>
              <Box>
                <Button
                  variant="contained"
                  disableRipple={true}
                  sx={{ ...styles.button, ...styles.primaryButton }}
                >
                  Reset Password
                </Button>
              </Box>
              <Box>
                <Button
                  variant="contained"
                  disableRipple={true}
                  href="/login"
                  sx={{
                    ...styles.button,
                    ...styles.primaryButton,
                    ...customStyle.backToLoginButton,
                  }}
                >
                  Back to Sign In
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

export default ForgotPage;
