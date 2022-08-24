import {
  Box,
  Checkbox,
  FormControlLabel,
  InputBase,
  NativeSelect,
  Typography,
} from "@mui/material";
import React from "react";

import { styles } from "../asset/style";

export interface IFormRowProps {
  rowHeader: string;
}

export const FormRow: React.FC<IFormRowProps> = ({ rowHeader }) => {
  return (
    <Box>
      <Typography variant="h1" sx={styles.inputLabel}>
        {rowHeader}
      </Typography>
      <Box component="form" sx={styles.inputBar}>
        <Box sx={styles.inputInner}>
          <InputBase sx={styles.inputBase} />
        </Box>
      </Box>
    </Box>
  );
};

export interface IFormCheckboxProps {
  label: string;
}

export const FormCheckbox: React.FC<IFormCheckboxProps> = ({ label }) => {
  return (
    <FormControlLabel
      control={<Checkbox />}
      label={
        <Typography variant="h1" sx={styles.inputLabel}>
          {label}
        </Typography>
      }
    />
  );
};

export interface IFormSelectProps {
  label: string;
  selectOptions: React.ReactNode;
}

export const FormSelect: React.FC<IFormSelectProps> = ({
  label,
  selectOptions,
}) => {
  return (
    <Box>
      <Typography variant="h1" sx={styles.inputLabel}>
        {label}
      </Typography>
      <Box component="form" sx={styles.inputBar}>
        <Box sx={styles.inputInner}>
          <NativeSelect disableUnderline={true}>{selectOptions}</NativeSelect>
        </Box>
      </Box>
    </Box>
  );
};
