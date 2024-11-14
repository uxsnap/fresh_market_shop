"use client";

import {
  DatesProvider,
  DateValue,
  DateInput as MantineDateInput,
} from "@mantine/dates";
import { ComponentProps, useEffect, useState } from "react";

import "@mantine/dates/styles.css";
import "dayjs/locale/ru";

import styles from "./DateInput.module.css";

export const DateInput = (props: ComponentProps<typeof MantineDateInput>) => {
  return (
    <DatesProvider settings={{ locale: "ru" }}>
      <MantineDateInput
        label="Дата Рождения"
        placeholder="ДД.ММ.ГГГГ"
        valueFormat="DD.MM.YYYY"
        classNames={{
          input: styles.input,
          weekdaysRow: styles.weekdaysRow,
          calendarHeader: styles.calendarHeader,
          day: styles.day,
        }}
        {...props}
        value={props.defaultValue}
      />
    </DatesProvider>
  );
};
