import {
  Box,
  Button,
  Card,
  CardContent,
  Link,
  Typography,
} from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import React from "react";

import { styles, theme } from "../../asset/style";
import Footer from "../../component/Footer";
import { FormRow } from "../../component/InputForm";
import TopNavBar from "../../component/TopNavBar";

const customStyle = {
  supportSpan: {
    alignItems: "center",
  },
  link: {
    display: "flex",
    alignItems: "flex-end",
  },
};

export interface ILoginPageProps {}

const LoginPage: React.FC<ILoginPageProps> = () => {
  return (
    <ThemeProvider theme={theme}>
      <div>
        <TopNavBar isAuth={false} />
        <Box sx={styles.container}>
          <Card sx={styles.loginForm}>
            <CardContent>
              <Typography variant="h1">Sign In</Typography>
              <span style={{ ...styles.container, ...customStyle.supportSpan }}>
                <Typography sx={styles.supportText}>
                  New to TutorKit?
                </Typography>
                <Button
                  disableRipple={true}
                  href="/register"
                  sx={{ ...styles.button, ...styles.secondaryButton }}
                >
                  Sign Up
                </Button>
              </span>
              <Box sx={styles.inputArea}>
                <FormRow rowHeader="Email" />
                <FormRow rowHeader="Password" />
              </Box>
              <Link href="/forgot" underline="none">
                <Typography
                  sx={{
                    ...styles.supportText,
                    ...customStyle.link,
                  }}
                >
                  Forgot your password?
                </Typography>
              </Link>
              <Box>
                <Button
                  variant="contained"
                  disableRipple={true}
                  sx={{ ...styles.button, ...styles.primaryButton }}
                >
                  Sign In
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

export default LoginPage;
