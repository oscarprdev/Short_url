import { API_URL } from '../../constants/apiUrl';
import WrapperAction from '../containers/WrapperAction';
import { IconBrandGithub, IconBrandGoogleFilled } from '@tabler/icons-react';

const DEFAULT_USER_URL = 'https://opr-short-url.vercel.app/user/116176187754032784002'

const LoginModal = () => {
	return (
		<article aria-label="login modal" className="w-[400px] flex flex-col items-center gap-10 pb-8">
			<h3 className="-mb-10">Choose an account</h3>
			<WrapperAction
				color="contrast-blue"
				url={DEFAULT_USER_URL}
				id="default"
				type="url">
				Default links
			</WrapperAction>
			<div className="flex items-center gap-5">
				<span className="w-[15px] h-[0.01rem] bg-primary" />
				<span className="w-[15px] h-[0.01rem] bg-primary" />
				<span className="w-[15px] h-[0.01rem] bg-primary" />
			</div>
			<div className="flex items-center gap-5 -mt-5">
				<WrapperAction color="default" url={`${API_URL}/auth?provider=google`} id="default" type="url">
					<IconBrandGoogleFilled />
					Google
				</WrapperAction>
				<WrapperAction color="default" url={`${API_URL}/auth?provider=github`} id="default" type="url">
					<IconBrandGithub />
					Github
				</WrapperAction>
			</div>
		</article>
	);
};

export default LoginModal;
