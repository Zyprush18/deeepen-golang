<script lang="ts">
    import * as Sidebar from "$lib/components/ui/sidebar/index.js";
    import AppSidebar from "$lib/components/layout/Sidebar-Layout.svelte";
    import Navbar from "$lib/components/layout/Navbar-Layout.svelte";

    let { children,profile,slug } = $props();
</script>

<Sidebar.Provider class="min-h-screen flex">
    <AppSidebar data={profile} />
    <main class="flex-1 block">
        <Navbar profil={profile} slug={slug} />
        {@render children?.()}
    </main>
</Sidebar.Provider>
