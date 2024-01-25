import { IconBrandGoogleFilled, IconBrandGithub, IconUser } from '@tabler/icons-react';
import LoginInput from '../ui/LoginInput';

const LoginModal = () => {
	return (
		<article className='w-[300px] flex flex-col items-center gap-5 pb-8'>
			<h3>Login with:</h3>
			<LoginInput
				href=''
				contrast>
				<IconUser />
				Default user
			</LoginInput>
			<div className='flex items-center gap-5'>
				<span className='w-[15px] h-[0.01rem] bg-primary' />
				<p>or</p>
				<span className='w-[15px] h-[0.01rem] bg-primary' />
			</div>
			<div className='flex items-center gap-5'>
				<LoginInput href='http://localhost:8080/auth?provider=google'>
					<IconBrandGoogleFilled />
					Google
				</LoginInput>
				<LoginInput href='http://localhost:8080/auth?provider=github'>
					<IconBrandGithub />
					Github
				</LoginInput>
			</div>
		</article>
	);
};

export default LoginModal;
