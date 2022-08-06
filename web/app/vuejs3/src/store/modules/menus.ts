import { defineStore } from "pinia";
import { api } from "@/http/api";
import { IMenu } from "../../types/menu";

export type RootState = {
  adminMenu: IMenu;
};

export const useMenusStore = defineStore("menus", {
  state: () =>
    ({
      adminMenu: {},
    } as RootState),
  getters: {
    admin: (state) => state.adminMenu,
  },
  actions: {
    async getAdminMenu() {
      return await api
        .get("http://localhost:9090/menus/1")
        .then((menu) => {
          this.adminMenu = menu.data;
        });
    },
  },
});
