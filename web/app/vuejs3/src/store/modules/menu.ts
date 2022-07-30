import { defineStore } from "pinia";
import { api } from "@/http/api";

export const useMenuStore = defineStore("menu", {
  state: () => ({
    menus: [],
  }),
  getters: {
    menusName: (state) => state.menus.map((el: { name: string }) => el.name),
  },
  actions: {
    async getMenus() {
      return await api
        .get("/menu?limit=20")
        .then((menus) => (this.menus = menus.data.results));
    },
  },
});
