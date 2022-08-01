<template>
    <component
               :is="iconComponent"
               role="img"
               class="inline-block fill-current uno-icon" />
</template>

<script>
import HomeIcon from "../icons/home.svg"
import UsersIcon from "../icons/users.svg"
import ContactsIcon from "../icons/contacts.svg"
import AccountsIcon from "../icons/accounts.svg"
import SettingsIcon from "../icons/settings.svg"
import HelpIcon from "../icons/help.svg"

const icons = {
    home: HomeIcon,
    users: UsersIcon,
    contacts: ContactsIcon,
    accounts: AccountsIcon,
    settings: SettingsIcon,
    help: HelpIcon
    // ...
}

export default {
    props: {
        name: {
            type: String,
            required: true,
            /*validator(value) {
                return Object.prototype.hasOwnProperty.call(icons, value)
            }*/
        }
    },

    computed: {
        iconComponent() {
            return icons[this.name]
        }
    }
}
</script>
<style scoped>
.uno-icon {
    width: 20px;
    height: 20px;
}
</style>