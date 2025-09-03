<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
</script>

<svelte:head>
    <title>Login</title>
</svelte:head>


<!-- <h1 class="text-5xl font-bold text-green-600 text-center">ZHAT</h1> -->
<div class="flex items-center justify-center h-screen">
    <Card.Root class="w-full max-w-sm">
        <Card.Header>
            <Card.Title class="text-center text-2xl">Login</Card.Title>
            <Card.Description>Enter your email below to login to your account</Card.Description>
        </Card.Header>
        <Card.Content>
            <form>
                <div class="flex flex-col gap-6 mb-3">
                    <div class="grid gap-2">
                        <Label for="email">Email</Label>
                        <Input id="email" type="email" placeholder="m@example.com" required />
                    </div>
                    <div class="grid gap-2">
                        <div class="flex items-center">
                            <Label for="password">Password</Label>
                            <a href="##" class="ml-auto inline-block text-sm underline-offset-4 hover:underline">
                                Forgot your password?
                            </a>
                        </div>
                        <Input id="password" type="password" required />
                    </div>
                </div>
                <span class="ml-auto inline-block text-sm underline-offset-4 text-gray-500 mb-2">
                        Belum Punya Akun? <a href="/register" class="text-blue-500 hover:underline">Register</a> 
                </span>
            </form>
        </Card.Content>
        <Card.Footer class="flex-col gap-2">
            <Button type="submit" class="w-full">Login</Button>
        </Card.Footer>
    </Card.Root>
</div>


