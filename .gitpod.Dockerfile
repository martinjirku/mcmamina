FROM gitpod/workspace-full

# Install Fly
RUN curl -L https://fly.io/install.sh | sh
ENV FLY_REGION="waw"
ENV FLYCTL_INSTALL="/home/gitpod/.fly"
ENV PATH="$FLYCTL_INSTALL/bin:$PATH"

# Install GitHub CLI
RUN brew install gh
