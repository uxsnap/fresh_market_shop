import { UserInfo } from "@/components/pages/profile/UserInfo";
import { Container, Group, Paper } from "@mantine/core";

export default function Profile() {
  return (
    <Container maw={1135} mx="auto" mt={36}>
      <Group gap={48} align="flex-start">
        <UserInfo />
      </Group>
    </Container>
  );
}
