<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import { actions } from './+page.server';

    

	console.log(actions.default)
</script>

<svelte:head>
	<title>Register</title>
</svelte:head>

<!-- <h1 class="text-5xl font-bold text-green-600 text-center">ZHAT</h1> -->
<div class="flex h-screen items-center justify-center">
	<Card.Root class="w-full max-w-sm">
		<Card.Header>
			<Card.Title class="text-center text-2xl">Register</Card.Title>
			<Card.Description>Enter your email below to create a new account</Card.Description>
		</Card.Header>
		<form method="POST">
			<Card.Content>
				<div class="mb-3 flex flex-col gap-6">
					<div class="grid gap-2">
						<Label for="username">Username</Label>
						<Input id="username" type="text" name="username" placeholder="Joko" required />
					</div>
					<div class="grid gap-2">
						<Label for="email">Email</Label>
						<Input id="email" type="email" name="email" placeholder="m@example.com" required />
					</div>
					<div class="grid gap-2">
						<div class="flex items-center">
							<Label for="password">Password</Label>
						</div>
						<Input id="password" type="password" name="password" placeholder="enter your password...." required />
					</div>
					<div class="grid gap-2">
						<div class="flex items-center">
							<Label for="repassword">Confirm Password</Label>
						</div>
						<Input
							id="repassword"
							type="password"
                            name="repassword"
							placeholder="enter your password again...."
							required
						/>
					</div>
				</div>
				<span class="ml-auto inline-block text-sm text-gray-500 underline-offset-4 mb-2">
					Sudah Punya Akun? <a href="/login" class="text-blue-500 hover:underline">Login</a>
				</span>
			</Card.Content>
			<Card.Footer class="flex-col gap-2">
				<Button type="submit" class="w-full">Register</Button>
			</Card.Footer>
		</form>
	</Card.Root>
</div>
