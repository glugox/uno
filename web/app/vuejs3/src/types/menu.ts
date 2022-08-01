import { ObjectId } from "./object";

export interface IMenu {
  id: ObjectId;
  label: string;
  items: IMenuItem[];
}

export interface IMenuItem {
  id: ObjectId;
  label: string;
  menu_id: string;
  parent_id: ObjectId;
  path: string;
  children: IMenuItem[];
}
