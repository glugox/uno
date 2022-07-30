import { reactive } from "vue";

export const navTree = reactive({
  nodes: [
    {
      id: "1",
      label: "Home",
    },
    {
      id: "2",
      label: "Users",
      children: [
        {
          id: "2-1",
          label: "List Users",
        },
        {
          id: "2-2",
          label: "Add User",
        },
      ],
    },
    {
      id: "3",
      label: "Contacts",
    },
    {
      id: "4",
      label: "Accounts",
    },
    {
      id: "5",
      label: "Settings",
    },
    {
      id: "6",
      label: "Help",
    },
  ],
  addItem() {
    this.nodes.push({
      id: "6",
      label: "Uno!",
    });
  },
});
