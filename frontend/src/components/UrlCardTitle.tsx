interface UrlCardTitleProps {
	title: string;
}
const UrlCardTitle = ({ title }: UrlCardTitleProps) => {
	return <p>{title}</p>;
};

export default UrlCardTitle;
