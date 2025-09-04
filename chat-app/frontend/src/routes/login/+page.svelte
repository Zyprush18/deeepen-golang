<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Card from '$lib/components/ui/card/index.js';


	let { form } = $props();
    console.log(form);
    

</script>

<svelte:head>
	<title>Login</title>
</svelte:head>

<!-- <h1 class="text-5xl font-bold text-green-600 text-center">ZHAT</h1> -->
<div class="flex h-screen items-center justify-center">
	<Card.Root class="w-full max-w-sm shadow-2xl">
		<Card.Header class="text-center">
			<span class="text-lg font-bold text-green-600">ZHAT</span>
			<Card.Title class="text-2xl">Login</Card.Title>
			<Card.Description>Enter your email below to login to your account</Card.Description>
		</Card.Header>
		<form method="POST">
			<Card.Content>
				<div class="mb-3 flex flex-col gap-6">
					<div class="grid gap-2">
						<Label for="email">Email</Label>
						<Input id="email" type="email" name="email" placeholder="m@example.com" />
					</div>
					<div class="grid gap-2">
						<div class="flex items-center">
							<Label for="password">Password</Label>
						</div>
						<Input
							id="password"
							type="password"
							name="password"
							placeholder="enter your password...."
						/>
					</div>
				</div>
				<span class="mb-5 ml-auto inline-block text-sm text-gray-500 underline-offset-4">
					Belum Punya Akun? <a href="/register" class="text-blue-500 hover:underline">Register</a>
				</span>
			</Card.Content>
			<Card.Footer class="flex-col gap-2">
				<Button type="submit" class="w-full bg-green-500 hover:bg-green-600">Login</Button>
			</Card.Footer>
		</form>
	</Card.Root>
</div>
