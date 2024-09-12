"use client";

import {
  DatesProvider,
  DatePickerInput as MantineDatePickerInput,
} from "@mantine/dates";
import { useState } from "react";

import "@mantine/dates/styles.css";
import "dayjs/locale/ru";

import styles from "./DateInput.module.css";

export const DateInput = () => {
  const [value, setValue] = useState<Date | null>(null);

  return (
    <DatesProvider settings={{ locale: "ru" }}>
      <MantineDatePickerInput
        value={value}
        onChange={setValue}
        label="Дата Рождения"
        placeholder="ДД.ММ.ГГГГ"
        classNames={{
          weekdaysRow: styles.weekdaysRow,
          calendarHeader: styles.calendarHeader,
          day: styles.day,
        }}
      />
    </DatesProvider>
  );
};
