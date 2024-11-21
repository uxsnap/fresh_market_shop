import { Addresses } from "@/components/pages/profile/Addresses";
import { Orders } from "@/components/pages/profile/Orders";
import { UserInfo } from "@/components/pages/profile/UserInfo";
import { Box, Group, Stack } from "@mantine/core";

import styles from "./profile.module.css";
import { ShadowBox } from "@/components/ShadowBox";

export default function Profile() {
  return (
    <Box maw={1135} mx="auto" mt={36}>
      <Group className={styles.wrapper}>
        <UserInfo />

        <Stack gap={16} w="100%">
          <ShadowBox w="100%" mah={312} style={{ overflowY: "auto" }}>
            <Addresses />
          </ShadowBox>

          <Orders />
        </Stack>
      </Group>
    </Box>
  );
}
