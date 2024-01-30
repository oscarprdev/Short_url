import { IconBrandGoogleFilled, IconBrandGithub, IconUser } from '@tabler/icons-react';
import AuthInput from '../AuthInput';
import { API_URL } from '../../constants/apiUrl';

const LoginModal = () => {
	return (
		<article className='w-[300px] flex flex-col items-center gap-5 pb-8'>
			<h3>Login with:</h3>
			<AuthInput
				href=''
				contrast>
				<IconUser />
				Default user
			</AuthInput>
			<div className='flex items-center gap-5'>
				<span className='w-[15px] h-[0.01rem] bg-primary' />
				<p>or</p>
				<span className='w-[15px] h-[0.01rem] bg-primary' />
			</div>
			<div className='flex items-center gap-5'>
				<AuthInput href={`${API_URL}/auth?provider=google`}>
					<IconBrandGoogleFilled />
					Google
				</AuthInput>
				<AuthInput href={`${API_URL}/auth?provider=github`}>
					<IconBrandGithub />
					Github
				</AuthInput>
			</div>
		</article>
	);
};

export default LoginModal;
