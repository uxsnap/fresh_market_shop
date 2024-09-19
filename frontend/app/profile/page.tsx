import { Addresses } from "@/components/pages/profile/Addresses";
import { Orders } from "@/components/pages/profile/Orders";
import { UserInfo } from "@/components/pages/profile/UserInfo";
import { Container, Group, Stack } from "@mantine/core";

export default function Profile() {
  return (
    <Container maw={1135} mx="auto" mt={36}>
      <Group wrap="nowrap" gap={48} align="flex-start">
        <UserInfo />

        <Stack gap={16} w="100%">
          <Addresses />

          <Orders />
        </Stack>
      </Group>
    </Container>
  );
}
