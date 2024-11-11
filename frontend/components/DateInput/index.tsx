"use client";

import {
  DatesProvider,
  DatePickerInput as MantineDatePickerInput,
} from "@mantine/dates";
import { ComponentProps } from "react";

import "@mantine/dates/styles.css";
import "dayjs/locale/ru";

import styles from "./DateInput.module.css";

export const DateInput = (
  props: ComponentProps<typeof MantineDatePickerInput>
) => (
  <DatesProvider settings={{ locale: "ru" }}>
    <MantineDatePickerInput
      label="Дата Рождения"
      placeholder="ДД.ММ.ГГГГ"
      classNames={{
        input: styles.input,
        weekdaysRow: styles.weekdaysRow,
        calendarHeader: styles.calendarHeader,
        day: styles.day,
      }}
      {...props}
    />
  </DatesProvider>
);
