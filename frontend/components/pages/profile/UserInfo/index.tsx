import { Avatar } from "@/components/Avatar";
import { DateInput } from "@/components/DateInput";
import { Select } from "@/components/Select";
import { ShadowBox } from "@/components/ShadowBox";
import { Box, Button, Group, Stack, TextInput } from "@mantine/core";

export const UserInfo = () => {
  return (
    <Box>
      <ShadowBox>
        <Stack gap={12}>
          <Box px={68} pt={20}>
            <Avatar />
          </Box>

          <Stack p={20} gap={12}>
            <TextInput size="md" label="Имя" placeholder="Введите имя" />

            <TextInput
              size="md"
              label="Фамилия"
              placeholder="Введите фамилию"
            />

            <DateInput />

            <Select
              label="Пол"
              placeholder="Выберите пол"
              data={["Мужской", "Женский"]}
            />
          </Stack>
        </Stack>
      </ShadowBox>

      <Group justify="space-between">
        <Button p={0} variant="outline">
          Выйти из системы
        </Button>
        <Button p={0} variant="outline" c="danger.0">
          Удалить аккаунт
        </Button>
      </Group>
    </Box>
  );
};
