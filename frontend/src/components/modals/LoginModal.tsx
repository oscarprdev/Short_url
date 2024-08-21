import { API_URL } from '../../constants/apiUrl';
import AuthInput from '../AuthInput';
import { IconBrandGithub, IconBrandGoogleFilled, IconUser } from '@tabler/icons-react';

// const DEFAULT_USER_URL = 'https://opr-short-url.vercel.app/user/116176187754032784002'

const LoginModal = () => {
	return (
		<article aria-label="login modal" className="w-[300px] flex flex-col items-center gap-5 pb-8">
			<h3>Login with:</h3>
			<AuthInput href={`http://localhost:5173/user/116176187754032784002`} contrast>
				<IconUser />
				Default user
			</AuthInput>
			<div className="flex items-center gap-5">
				<span className="w-[15px] h-[0.01rem] bg-primary" />
				<p>or</p>
				<span className="w-[15px] h-[0.01rem] bg-primary" />
			</div>
			<div className="flex items-center gap-5">
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
