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
        .get("http://localhost:9090/menus/25A9B312D0440797C2995B1E3784C488")
        .then((menu) => {
          this.adminMenu = menu.data;
        });
    },
  },
});
